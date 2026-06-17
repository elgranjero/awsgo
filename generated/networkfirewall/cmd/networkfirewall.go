package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// networkfirewallCmd represents the networkfirewall command
var _networkfirewallCmd = &cobra.Command{
	Use:   "networkfirewall",
	Short: "AWS networkfirewall CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := networkfirewall.NewFromConfig(cfg)
		if _networkfirewallAcceptNetworkFirewallTransitGatewayAttachment {
			networkfirewall_AcceptNetworkFirewallTransitGatewayAttachment(cfg, client)
			return
		}
		if _networkfirewallAssociateAvailabilityZones {
			networkfirewall_AssociateAvailabilityZones(cfg, client)
			return
		}
		if _networkfirewallAssociateFirewallPolicy {
			networkfirewall_AssociateFirewallPolicy(cfg, client)
			return
		}
		if _networkfirewallAssociateSubnets {
			networkfirewall_AssociateSubnets(cfg, client)
			return
		}
		if _networkfirewallAttachRuleGroupsToProxyConfiguration {
			networkfirewall_AttachRuleGroupsToProxyConfiguration(cfg, client)
			return
		}
		if _networkfirewallCreateFirewall {
			networkfirewall_CreateFirewall(cfg, client)
			return
		}
		if _networkfirewallCreateFirewallPolicy {
			networkfirewall_CreateFirewallPolicy(cfg, client)
			return
		}
		if _networkfirewallCreateProxy {
			networkfirewall_CreateProxy(cfg, client)
			return
		}
		if _networkfirewallCreateProxyConfiguration {
			networkfirewall_CreateProxyConfiguration(cfg, client)
			return
		}
		if _networkfirewallCreateProxyRuleGroup {
			networkfirewall_CreateProxyRuleGroup(cfg, client)
			return
		}
		if _networkfirewallCreateProxyRules {
			networkfirewall_CreateProxyRules(cfg, client)
			return
		}
		if _networkfirewallCreateRuleGroup {
			networkfirewall_CreateRuleGroup(cfg, client)
			return
		}
		if _networkfirewallCreateTLSInspectionConfiguration {
			networkfirewall_CreateTLSInspectionConfiguration(cfg, client)
			return
		}
		if _networkfirewallCreateVpcEndpointAssociation {
			networkfirewall_CreateVpcEndpointAssociation(cfg, client)
			return
		}
		if _networkfirewallDeleteFirewall {
			networkfirewall_DeleteFirewall(cfg, client)
			return
		}
		if _networkfirewallDeleteFirewallPolicy {
			networkfirewall_DeleteFirewallPolicy(cfg, client)
			return
		}
		if _networkfirewallDeleteNetworkFirewallTransitGatewayAttachment {
			networkfirewall_DeleteNetworkFirewallTransitGatewayAttachment(cfg, client)
			return
		}
		if _networkfirewallDeleteProxy {
			networkfirewall_DeleteProxy(cfg, client)
			return
		}
		if _networkfirewallDeleteProxyConfiguration {
			networkfirewall_DeleteProxyConfiguration(cfg, client)
			return
		}
		if _networkfirewallDeleteProxyRuleGroup {
			networkfirewall_DeleteProxyRuleGroup(cfg, client)
			return
		}
		if _networkfirewallDeleteProxyRules {
			networkfirewall_DeleteProxyRules(cfg, client)
			return
		}
		if _networkfirewallDeleteResourcePolicy {
			networkfirewall_DeleteResourcePolicy(cfg, client)
			return
		}
		if _networkfirewallDeleteRuleGroup {
			networkfirewall_DeleteRuleGroup(cfg, client)
			return
		}
		if _networkfirewallDeleteTLSInspectionConfiguration {
			networkfirewall_DeleteTLSInspectionConfiguration(cfg, client)
			return
		}
		if _networkfirewallDeleteVpcEndpointAssociation {
			networkfirewall_DeleteVpcEndpointAssociation(cfg, client)
			return
		}
		if _networkfirewallDescribeFirewall {
			networkfirewall_DescribeFirewall(cfg, client)
			return
		}
		if _networkfirewallDescribeFirewallMetadata {
			networkfirewall_DescribeFirewallMetadata(cfg, client)
			return
		}
		if _networkfirewallDescribeFirewallPolicy {
			networkfirewall_DescribeFirewallPolicy(cfg, client)
			return
		}
		if _networkfirewallDescribeFlowOperation {
			networkfirewall_DescribeFlowOperation(cfg, client)
			return
		}
		if _networkfirewallDescribeLoggingConfiguration {
			networkfirewall_DescribeLoggingConfiguration(cfg, client)
			return
		}
		if _networkfirewallDescribeProxy {
			networkfirewall_DescribeProxy(cfg, client)
			return
		}
		if _networkfirewallDescribeProxyConfiguration {
			networkfirewall_DescribeProxyConfiguration(cfg, client)
			return
		}
		if _networkfirewallDescribeProxyRule {
			networkfirewall_DescribeProxyRule(cfg, client)
			return
		}
		if _networkfirewallDescribeProxyRuleGroup {
			networkfirewall_DescribeProxyRuleGroup(cfg, client)
			return
		}
		if _networkfirewallDescribeResourcePolicy {
			networkfirewall_DescribeResourcePolicy(cfg, client)
			return
		}
		if _networkfirewallDescribeRuleGroup {
			networkfirewall_DescribeRuleGroup(cfg, client)
			return
		}
		if _networkfirewallDescribeRuleGroupMetadata {
			networkfirewall_DescribeRuleGroupMetadata(cfg, client)
			return
		}
		if _networkfirewallDescribeRuleGroupSummary {
			networkfirewall_DescribeRuleGroupSummary(cfg, client)
			return
		}
		if _networkfirewallDescribeTLSInspectionConfiguration {
			networkfirewall_DescribeTLSInspectionConfiguration(cfg, client)
			return
		}
		if _networkfirewallDescribeVpcEndpointAssociation {
			networkfirewall_DescribeVpcEndpointAssociation(cfg, client)
			return
		}
		if _networkfirewallDetachRuleGroupsFromProxyConfiguration {
			networkfirewall_DetachRuleGroupsFromProxyConfiguration(cfg, client)
			return
		}
		if _networkfirewallDisassociateAvailabilityZones {
			networkfirewall_DisassociateAvailabilityZones(cfg, client)
			return
		}
		if _networkfirewallDisassociateSubnets {
			networkfirewall_DisassociateSubnets(cfg, client)
			return
		}
		if _networkfirewallGetAnalysisReportResults {
			networkfirewall_GetAnalysisReportResults(cfg, client)
			return
		}
		if _networkfirewallListAnalysisReports {
			networkfirewall_ListAnalysisReports(cfg, client)
			return
		}
		if _networkfirewallListFirewallPolicies {
			networkfirewall_ListFirewallPolicies(cfg, client)
			return
		}
		if _networkfirewallListFirewalls {
			networkfirewall_ListFirewalls(cfg, client)
			return
		}
		if _networkfirewallListFlowOperationResults {
			networkfirewall_ListFlowOperationResults(cfg, client)
			return
		}
		if _networkfirewallListFlowOperations {
			networkfirewall_ListFlowOperations(cfg, client)
			return
		}
		if _networkfirewallListProxies {
			networkfirewall_ListProxies(cfg, client)
			return
		}
		if _networkfirewallListProxyConfigurations {
			networkfirewall_ListProxyConfigurations(cfg, client)
			return
		}
		if _networkfirewallListProxyRuleGroups {
			networkfirewall_ListProxyRuleGroups(cfg, client)
			return
		}
		if _networkfirewallListRuleGroups {
			networkfirewall_ListRuleGroups(cfg, client)
			return
		}
		if _networkfirewallListTagsForResource {
			networkfirewall_ListTagsForResource(cfg, client)
			return
		}
		if _networkfirewallListTLSInspectionConfigurations {
			networkfirewall_ListTLSInspectionConfigurations(cfg, client)
			return
		}
		if _networkfirewallListVpcEndpointAssociations {
			networkfirewall_ListVpcEndpointAssociations(cfg, client)
			return
		}
		if _networkfirewallPutResourcePolicy {
			networkfirewall_PutResourcePolicy(cfg, client)
			return
		}
		if _networkfirewallRejectNetworkFirewallTransitGatewayAttachment {
			networkfirewall_RejectNetworkFirewallTransitGatewayAttachment(cfg, client)
			return
		}
		if _networkfirewallStartAnalysisReport {
			networkfirewall_StartAnalysisReport(cfg, client)
			return
		}
		if _networkfirewallStartFlowCapture {
			networkfirewall_StartFlowCapture(cfg, client)
			return
		}
		if _networkfirewallStartFlowFlush {
			networkfirewall_StartFlowFlush(cfg, client)
			return
		}
		if _networkfirewallTagResource {
			networkfirewall_TagResource(cfg, client)
			return
		}
		if _networkfirewallUntagResource {
			networkfirewall_UntagResource(cfg, client)
			return
		}
		if _networkfirewallUpdateAvailabilityZoneChangeProtection {
			networkfirewall_UpdateAvailabilityZoneChangeProtection(cfg, client)
			return
		}
		if _networkfirewallUpdateFirewallAnalysisSettings {
			networkfirewall_UpdateFirewallAnalysisSettings(cfg, client)
			return
		}
		if _networkfirewallUpdateFirewallDeleteProtection {
			networkfirewall_UpdateFirewallDeleteProtection(cfg, client)
			return
		}
		if _networkfirewallUpdateFirewallDescription {
			networkfirewall_UpdateFirewallDescription(cfg, client)
			return
		}
		if _networkfirewallUpdateFirewallEncryptionConfiguration {
			networkfirewall_UpdateFirewallEncryptionConfiguration(cfg, client)
			return
		}
		if _networkfirewallUpdateFirewallPolicy {
			networkfirewall_UpdateFirewallPolicy(cfg, client)
			return
		}
		if _networkfirewallUpdateFirewallPolicyChangeProtection {
			networkfirewall_UpdateFirewallPolicyChangeProtection(cfg, client)
			return
		}
		if _networkfirewallUpdateLoggingConfiguration {
			networkfirewall_UpdateLoggingConfiguration(cfg, client)
			return
		}
		if _networkfirewallUpdateProxy {
			networkfirewall_UpdateProxy(cfg, client)
			return
		}
		if _networkfirewallUpdateProxyConfiguration {
			networkfirewall_UpdateProxyConfiguration(cfg, client)
			return
		}
		if _networkfirewallUpdateProxyRule {
			networkfirewall_UpdateProxyRule(cfg, client)
			return
		}
		if _networkfirewallUpdateProxyRuleGroupPriorities {
			networkfirewall_UpdateProxyRuleGroupPriorities(cfg, client)
			return
		}
		if _networkfirewallUpdateProxyRulePriorities {
			networkfirewall_UpdateProxyRulePriorities(cfg, client)
			return
		}
		if _networkfirewallUpdateRuleGroup {
			networkfirewall_UpdateRuleGroup(cfg, client)
			return
		}
		if _networkfirewallUpdateSubnetChangeProtection {
			networkfirewall_UpdateSubnetChangeProtection(cfg, client)
			return
		}
		if _networkfirewallUpdateTLSInspectionConfiguration {
			networkfirewall_UpdateTLSInspectionConfiguration(cfg, client)
			return
		}

	},
}

var (
	_networkfirewallAcceptNetworkFirewallTransitGatewayAttachment bool
	_networkfirewallAssociateAvailabilityZones                    bool
	_networkfirewallAssociateFirewallPolicy                       bool
	_networkfirewallAssociateSubnets                              bool
	_networkfirewallAttachRuleGroupsToProxyConfiguration          bool
	_networkfirewallCreateFirewall                                bool
	_networkfirewallCreateFirewallPolicy                          bool
	_networkfirewallCreateProxy                                   bool
	_networkfirewallCreateProxyConfiguration                      bool
	_networkfirewallCreateProxyRuleGroup                          bool
	_networkfirewallCreateProxyRules                              bool
	_networkfirewallCreateRuleGroup                               bool
	_networkfirewallCreateTLSInspectionConfiguration              bool
	_networkfirewallCreateVpcEndpointAssociation                  bool
	_networkfirewallDeleteFirewall                                bool
	_networkfirewallDeleteFirewallPolicy                          bool
	_networkfirewallDeleteNetworkFirewallTransitGatewayAttachment bool
	_networkfirewallDeleteProxy                                   bool
	_networkfirewallDeleteProxyConfiguration                      bool
	_networkfirewallDeleteProxyRuleGroup                          bool
	_networkfirewallDeleteProxyRules                              bool
	_networkfirewallDeleteResourcePolicy                          bool
	_networkfirewallDeleteRuleGroup                               bool
	_networkfirewallDeleteTLSInspectionConfiguration              bool
	_networkfirewallDeleteVpcEndpointAssociation                  bool
	_networkfirewallDescribeFirewall                              bool
	_networkfirewallDescribeFirewallMetadata                      bool
	_networkfirewallDescribeFirewallPolicy                        bool
	_networkfirewallDescribeFlowOperation                         bool
	_networkfirewallDescribeLoggingConfiguration                  bool
	_networkfirewallDescribeProxy                                 bool
	_networkfirewallDescribeProxyConfiguration                    bool
	_networkfirewallDescribeProxyRule                             bool
	_networkfirewallDescribeProxyRuleGroup                        bool
	_networkfirewallDescribeResourcePolicy                        bool
	_networkfirewallDescribeRuleGroup                             bool
	_networkfirewallDescribeRuleGroupMetadata                     bool
	_networkfirewallDescribeRuleGroupSummary                      bool
	_networkfirewallDescribeTLSInspectionConfiguration            bool
	_networkfirewallDescribeVpcEndpointAssociation                bool
	_networkfirewallDetachRuleGroupsFromProxyConfiguration        bool
	_networkfirewallDisassociateAvailabilityZones                 bool
	_networkfirewallDisassociateSubnets                           bool
	_networkfirewallGetAnalysisReportResults                      bool
	_networkfirewallListAnalysisReports                           bool
	_networkfirewallListFirewallPolicies                          bool
	_networkfirewallListFirewalls                                 bool
	_networkfirewallListFlowOperationResults                      bool
	_networkfirewallListFlowOperations                            bool
	_networkfirewallListProxies                                   bool
	_networkfirewallListProxyConfigurations                       bool
	_networkfirewallListProxyRuleGroups                           bool
	_networkfirewallListRuleGroups                                bool
	_networkfirewallListTagsForResource                           bool
	_networkfirewallListTLSInspectionConfigurations               bool
	_networkfirewallListVpcEndpointAssociations                   bool
	_networkfirewallPutResourcePolicy                             bool
	_networkfirewallRejectNetworkFirewallTransitGatewayAttachment bool
	_networkfirewallStartAnalysisReport                           bool
	_networkfirewallStartFlowCapture                              bool
	_networkfirewallStartFlowFlush                                bool
	_networkfirewallTagResource                                   bool
	_networkfirewallUntagResource                                 bool
	_networkfirewallUpdateAvailabilityZoneChangeProtection        bool
	_networkfirewallUpdateFirewallAnalysisSettings                bool
	_networkfirewallUpdateFirewallDeleteProtection                bool
	_networkfirewallUpdateFirewallDescription                     bool
	_networkfirewallUpdateFirewallEncryptionConfiguration         bool
	_networkfirewallUpdateFirewallPolicy                          bool
	_networkfirewallUpdateFirewallPolicyChangeProtection          bool
	_networkfirewallUpdateLoggingConfiguration                    bool
	_networkfirewallUpdateProxy                                   bool
	_networkfirewallUpdateProxyConfiguration                      bool
	_networkfirewallUpdateProxyRule                               bool
	_networkfirewallUpdateProxyRuleGroupPriorities                bool
	_networkfirewallUpdateProxyRulePriorities                     bool
	_networkfirewallUpdateRuleGroup                               bool
	_networkfirewallUpdateSubnetChangeProtection                  bool
	_networkfirewallUpdateTLSInspectionConfiguration              bool

	_networkfirewallAction                           string
	_networkfirewallAddConditions                    string
	_networkfirewallAnalysisReportId                 string
	_networkfirewallAnalysisType                     string
	_networkfirewallAnalyzeRuleGroup                 string
	_networkfirewallAvailabilityZone                 string
	_networkfirewallAvailabilityZoneChangeProtection string
	_networkfirewallAvailabilityZoneMappings         string
	_networkfirewallCapacity                         string
	_networkfirewallDefaultRulePhaseActions          string
	_networkfirewallDeleteProtection                 string
	_networkfirewallDescription                      string
	_networkfirewallDryRun                           string
	_networkfirewallEnableMonitoringDashboard        string
	_networkfirewallEnabledAnalysisTypes             string
	_networkfirewallEncryptionConfiguration          string
	_networkfirewallFirewallArn                      string
	_networkfirewallFirewallName                     string
	_networkfirewallFirewallPolicy                   string
	_networkfirewallFirewallPolicyArn                string
	_networkfirewallFirewallPolicyChangeProtection   string
	_networkfirewallFirewallPolicyName               string
	_networkfirewallFlowFilters                      string
	_networkfirewallFlowOperationId                  string
	_networkfirewallFlowOperationType                string
	_networkfirewallListenerProperties               string
	_networkfirewallListenerPropertiesToAdd          string
	_networkfirewallListenerPropertiesToRemove       string
	_networkfirewallLoggingConfiguration             string
	_networkfirewallManagedType                      string
	_networkfirewallMaxResults                       string
	_networkfirewallMinimumFlowAgeInSeconds          string
	_networkfirewallNatGatewayId                     string
	_networkfirewallNextToken                        string
	_networkfirewallPolicy                           string
	_networkfirewallProxyArn                         string
	_networkfirewallProxyConfigurationArn            string
	_networkfirewallProxyConfigurationName           string
	_networkfirewallProxyName                        string
	_networkfirewallProxyRuleGroupArn                string
	_networkfirewallProxyRuleGroupName               string
	_networkfirewallProxyRuleName                    string
	_networkfirewallRemoveConditions                 string
	_networkfirewallResourceArn                      string
	_networkfirewallRuleGroup                        string
	_networkfirewallRuleGroupArn                     string
	_networkfirewallRuleGroupArns                    []string
	_networkfirewallRuleGroupName                    string
	_networkfirewallRuleGroupNames                   []string
	_networkfirewallRuleGroupRequestPhase            string
	_networkfirewallRuleGroups                       string
	_networkfirewallRules                            string
	_networkfirewallScope                            string
	_networkfirewallSourceMetadata                   string
	_networkfirewallSubnetChangeProtection           string
	_networkfirewallSubnetIds                        []string
	_networkfirewallSubnetMapping                    string
	_networkfirewallSubnetMappings                   string
	_networkfirewallSubscriptionStatus               string
	_networkfirewallSummaryConfiguration             string
	_networkfirewallTagKeys                          []string
	_networkfirewallTags                             string
	_networkfirewallTLSInspectionConfiguration       string
	_networkfirewallTLSInspectionConfigurationArn    string
	_networkfirewallTLSInspectionConfigurationName   string
	_networkfirewallTlsInterceptProperties           string
	_networkfirewallTransitGatewayAttachmentId       string
	_networkfirewallTransitGatewayId                 string
	_networkfirewallType                             string
	_networkfirewallUpdateToken                      string
	_networkfirewallVpcEndpointAssociationArn        string
	_networkfirewallVpcEndpointId                    string
	_networkfirewallVpcId                            string
	_networkfirewallVpcIds                           []string
)

// Accepts a transit gateway attachment request for Network Firewall. When you
// accept the attachment request, Network Firewall creates the necessary routing
// components to enable traffic flow between the transit gateway and firewall
// endpoints.
//
// You must accept a transit gateway attachment to complete the creation of a
// transit gateway-attached firewall, unless auto-accept is enabled on the transit
// gateway. After acceptance, use DescribeFirewallto verify the firewall status.
//
// To reject an attachment instead of accepting it, use RejectNetworkFirewallTransitGatewayAttachment.
//
// It can take several minutes for the attachment acceptance to complete and the
// firewall to become available.
func networkfirewall_AcceptNetworkFirewallTransitGatewayAttachment(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.AcceptNetworkFirewallTransitGatewayAttachmentInput{
		// TransitGatewayAttachmentId: *string, // Required
	}

	if len(_networkfirewallTransitGatewayAttachmentId) > 0 {
		input.TransitGatewayAttachmentId = aws.String(_networkfirewallTransitGatewayAttachmentId)
	}

	if resp, err := client.AcceptNetworkFirewallTransitGatewayAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified Availability Zones with a transit gateway-attached
// firewall. For each Availability Zone, Network Firewall creates a firewall
// endpoint to process traffic. You can specify one or more Availability Zones
// where you want to deploy the firewall.
//
// After adding Availability Zones, you must update your transit gateway route
// tables to direct traffic through the new firewall endpoints. Use DescribeFirewallto monitor the
// status of the new endpoints.
func networkfirewall_AssociateAvailabilityZones(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.AssociateAvailabilityZonesInput{
		// AvailabilityZoneMappings: []types.AvailabilityZoneMapping, // Required
	}

	if len(_networkfirewallAvailabilityZoneMappings) > 0 {
		if err := assignInputField(input, "AvailabilityZoneMappings", _networkfirewallAvailabilityZoneMappings); err != nil {
			log.Errorf("invalid --availability-zone-mappings: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.AssociateAvailabilityZones(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a FirewallPolicy to a Firewall.
// A firewall policy defines how to monitor and manage your VPC network traffic,
// using a collection of inspection rule groups and other settings. Each firewall
// requires one firewall policy association, and you can use the same firewall
// policy for multiple firewalls.
func networkfirewall_AssociateFirewallPolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.AssociateFirewallPolicyInput{
		// FirewallPolicyArn: *string, // Required
	}

	if len(_networkfirewallFirewallPolicyArn) > 0 {
		input.FirewallPolicyArn = aws.String(_networkfirewallFirewallPolicyArn)
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.AssociateFirewallPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified subnets in the Amazon VPC to the firewall. You can
// specify one subnet for each of the Availability Zones that the VPC spans.
//
// This request creates an Network Firewall firewall endpoint in each of the
// subnets. To enable the firewall's protections, you must also modify the VPC's
// route tables for each subnet's Availability Zone, to redirect the traffic that's
// coming into and going out of the zone through the firewall endpoint.
func networkfirewall_AssociateSubnets(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.AssociateSubnetsInput{
		// SubnetMappings: []types.SubnetMapping, // Required
	}

	if len(_networkfirewallSubnetMappings) > 0 {
		if err := assignInputField(input, "SubnetMappings", _networkfirewallSubnetMappings); err != nil {
			log.Errorf("invalid --subnet-mappings: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.AssociateSubnets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches ProxyRuleGroup resources to a ProxyConfiguration
// A Proxy Configuration defines the monitoring and protection behavior for a
// Proxy. The details of the behavior are defined in the rule groups that you add
// to your configuration.
func networkfirewall_AttachRuleGroupsToProxyConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.AttachRuleGroupsToProxyConfigurationInput{
		// RuleGroups: []types.ProxyRuleGroupAttachment, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallRuleGroups) > 0 {
		if err := assignInputField(input, "RuleGroups", _networkfirewallRuleGroups); err != nil {
			log.Errorf("invalid --rule-groups: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}

	if resp, err := client.AttachRuleGroupsToProxyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Network Firewall Firewall and accompanying FirewallStatus for a VPC.
// The firewall defines the configuration settings for an Network Firewall
// firewall. The settings that you can define at creation include the firewall
// policy, the subnets in your VPC to use for the firewall endpoints, and any tags
// that are attached to the firewall Amazon Web Services resource.
//
// After you create a firewall, you can provide additional settings, like the
// logging configuration.
//
// To update the settings for a firewall, you use the operations that apply to the
// settings themselves, for example UpdateLoggingConfiguration, AssociateSubnets, and UpdateFirewallDeleteProtection.
//
// To manage a firewall's tags, use the standard Amazon Web Services resource
// tagging operations, ListTagsForResource, TagResource, and UntagResource.
//
// To retrieve information about firewalls, use ListFirewalls and DescribeFirewall.
//
// To generate a report on the last 30 days of traffic monitored by a firewall,
// use StartAnalysisReport.
func networkfirewall_CreateFirewall(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateFirewallInput{
		// FirewallName: *string, // Required
		// FirewallPolicyArn: *string, // Required
	}

	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallFirewallPolicyArn) > 0 {
		input.FirewallPolicyArn = aws.String(_networkfirewallFirewallPolicyArn)
	}
	if len(_networkfirewallAvailabilityZoneChangeProtection) > 0 {
		if err := assignInputField(input, "AvailabilityZoneChangeProtection", _networkfirewallAvailabilityZoneChangeProtection); err != nil {
			log.Errorf("invalid --availability-zone-change-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallAvailabilityZoneMappings) > 0 {
		if err := assignInputField(input, "AvailabilityZoneMappings", _networkfirewallAvailabilityZoneMappings); err != nil {
			log.Errorf("invalid --availability-zone-mappings: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallDeleteProtection) > 0 {
		if err := assignInputField(input, "DeleteProtection", _networkfirewallDeleteProtection); err != nil {
			log.Errorf("invalid --delete-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallEnabledAnalysisTypes) > 0 {
		if err := assignInputField(input, "EnabledAnalysisTypes", _networkfirewallEnabledAnalysisTypes); err != nil {
			log.Errorf("invalid --enabled-analysis-types: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallPolicyChangeProtection) > 0 {
		if err := assignInputField(input, "FirewallPolicyChangeProtection", _networkfirewallFirewallPolicyChangeProtection); err != nil {
			log.Errorf("invalid --firewall-policy-change-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallSubnetChangeProtection) > 0 {
		if err := assignInputField(input, "SubnetChangeProtection", _networkfirewallSubnetChangeProtection); err != nil {
			log.Errorf("invalid --subnet-change-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallSubnetMappings) > 0 {
		if err := assignInputField(input, "SubnetMappings", _networkfirewallSubnetMappings); err != nil {
			log.Errorf("invalid --subnet-mappings: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTransitGatewayId) > 0 {
		input.TransitGatewayId = aws.String(_networkfirewallTransitGatewayId)
	}
	if len(_networkfirewallVpcId) > 0 {
		input.VpcId = aws.String(_networkfirewallVpcId)
	}

	if resp, err := client.CreateFirewall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the firewall policy for the firewall according to the specifications.
// An Network Firewall firewall policy defines the behavior of a firewall, in a
// collection of stateless and stateful rule groups and other settings. You can use
// one firewall policy for multiple firewalls.
func networkfirewall_CreateFirewallPolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateFirewallPolicyInput{
		// FirewallPolicy: *types.FirewallPolicy, // Required
		// FirewallPolicyName: *string, // Required
	}

	if len(_networkfirewallFirewallPolicy) > 0 {
		if err := assignInputField(input, "FirewallPolicy", _networkfirewallFirewallPolicy); err != nil {
			log.Errorf("invalid --firewall-policy: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallPolicyName) > 0 {
		input.FirewallPolicyName = aws.String(_networkfirewallFirewallPolicyName)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _networkfirewallDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFirewallPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Network Firewall Proxy
// Attaches a Proxy configuration to a NAT Gateway.
//
// To manage a proxy's tags, use the standard Amazon Web Services resource tagging
// operations, ListTagsForResource, TagResource, and UntagResource.
//
// To retrieve information about proxies, use ListProxies and DescribeProxy.
func networkfirewall_CreateProxy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateProxyInput{
		// NatGatewayId: *string, // Required
		// ProxyName: *string, // Required
		// TlsInterceptProperties: *types.TlsInterceptPropertiesRequest, // Required
	}

	if len(_networkfirewallNatGatewayId) > 0 {
		input.NatGatewayId = aws.String(_networkfirewallNatGatewayId)
	}
	if len(_networkfirewallProxyName) > 0 {
		input.ProxyName = aws.String(_networkfirewallProxyName)
	}
	if len(_networkfirewallTlsInterceptProperties) > 0 {
		if err := assignInputField(input, "TlsInterceptProperties", _networkfirewallTlsInterceptProperties); err != nil {
			log.Errorf("invalid --tls-intercept-properties: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallListenerProperties) > 0 {
		if err := assignInputField(input, "ListenerProperties", _networkfirewallListenerProperties); err != nil {
			log.Errorf("invalid --listener-properties: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Network Firewall ProxyConfiguration
// A Proxy Configuration defines the monitoring and protection behavior for a
// Proxy. The details of the behavior are defined in the rule groups that you add
// to your configuration.
//
// To manage a proxy configuration's tags, use the standard Amazon Web Services
// resource tagging operations, ListTagsForResource, TagResource, and UntagResource.
//
// To retrieve information about proxies, use ListProxyConfigurations and DescribeProxyConfiguration.
func networkfirewall_CreateProxyConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateProxyConfigurationInput{
		// DefaultRulePhaseActions: *types.ProxyConfigDefaultRulePhaseActionsRequest, // Required
		// ProxyConfigurationName: *string, // Required
	}

	if len(_networkfirewallDefaultRulePhaseActions) > 0 {
		if err := assignInputField(input, "DefaultRulePhaseActions", _networkfirewallDefaultRulePhaseActions); err != nil {
			log.Errorf("invalid --default-rule-phase-actions: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallRuleGroupArns) > 0 {
		input.RuleGroupArns = append([]string(nil), _networkfirewallRuleGroupArns...)
	}
	if len(_networkfirewallRuleGroupNames) > 0 {
		input.RuleGroupNames = append([]string(nil), _networkfirewallRuleGroupNames...)
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProxyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Network Firewall ProxyRuleGroup
// Collections of related proxy filtering rules. Rule groups help you manage and
// reuse sets of rules across multiple proxy configurations.
//
// To manage a proxy rule group's tags, use the standard Amazon Web Services
// resource tagging operations, ListTagsForResource, TagResource, and UntagResource.
//
// To retrieve information about proxy rule groups, use ListProxyRuleGroups and DescribeProxyRuleGroup.
//
// To retrieve information about individual proxy rules, use DescribeProxyRuleGroup and DescribeProxyRule.
func networkfirewall_CreateProxyRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateProxyRuleGroupInput{
		// ProxyRuleGroupName: *string, // Required
	}

	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallRules) > 0 {
		if err := assignInputField(input, "Rules", _networkfirewallRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProxyRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates Network Firewall ProxyRule resources.
// Attaches new proxy rule(s) to an existing proxy rule group.
//
// To retrieve information about individual proxy rules, use DescribeProxyRuleGroup and DescribeProxyRule.
func networkfirewall_CreateProxyRules(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateProxyRulesInput{
		// Rules: *types.CreateProxyRulesByRequestPhase, // Required
	}

	if len(_networkfirewallRules) > 0 {
		if err := assignInputField(input, "Rules", _networkfirewallRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}

	if resp, err := client.CreateProxyRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified stateless or stateful rule group, which includes the
// rules for network traffic inspection, a capacity setting, and tags.
//
// You provide your rule group specification in your request using either RuleGroup
// or Rules .
func networkfirewall_CreateRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateRuleGroupInput{
		// Capacity: *int32, // Required
		// RuleGroupName: *string, // Required
		// Type: types.RuleGroupType, // Required
	}

	if len(_networkfirewallCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _networkfirewallCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRuleGroupName) > 0 {
		input.RuleGroupName = aws.String(_networkfirewallRuleGroupName)
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallAnalyzeRuleGroup) > 0 {
		if err := assignInputField(input, "AnalyzeRuleGroup", _networkfirewallAnalyzeRuleGroup); err != nil {
			log.Errorf("invalid --analyze-rule-group: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _networkfirewallDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRuleGroup) > 0 {
		if err := assignInputField(input, "RuleGroup", _networkfirewallRuleGroup); err != nil {
			log.Errorf("invalid --rule-group: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRules) > 0 {
		input.Rules = aws.String(_networkfirewallRules)
	}
	if len(_networkfirewallSourceMetadata) > 0 {
		if err := assignInputField(input, "SourceMetadata", _networkfirewallSourceMetadata); err != nil {
			log.Errorf("invalid --source-metadata: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallSummaryConfiguration) > 0 {
		if err := assignInputField(input, "SummaryConfiguration", _networkfirewallSummaryConfiguration); err != nil {
			log.Errorf("invalid --summary-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Network Firewall TLS inspection configuration. Network Firewall uses
// TLS inspection configurations to decrypt your firewall's inbound and outbound
// SSL/TLS traffic. After decryption, Network Firewall inspects the traffic
// according to your firewall policy's stateful rules, and then re-encrypts it
// before sending it to its destination. You can enable inspection of your
// firewall's inbound traffic, outbound traffic, or both. To use TLS inspection
// with your firewall, you must first import or provision certificates using ACM,
// create a TLS inspection configuration, add that configuration to a new firewall
// policy, and then associate that policy with your firewall.
//
// To update the settings for a TLS inspection configuration, use UpdateTLSInspectionConfiguration.
//
// To manage a TLS inspection configuration's tags, use the standard Amazon Web
// Services resource tagging operations, ListTagsForResource, TagResource, and UntagResource.
//
// To retrieve information about TLS inspection configurations, use ListTLSInspectionConfigurations and DescribeTLSInspectionConfiguration.
//
// For more information about TLS inspection configurations, see [Inspecting SSL/TLS traffic with TLS inspection configurations] in the Network
// Firewall Developer Guide.
//
// [Inspecting SSL/TLS traffic with TLS inspection configurations]: https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html
func networkfirewall_CreateTLSInspectionConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateTLSInspectionConfigurationInput{
		// TLSInspectionConfiguration: *types.TLSInspectionConfiguration, // Required
		// TLSInspectionConfigurationName: *string, // Required
	}

	if len(_networkfirewallTLSInspectionConfiguration) > 0 {
		if err := assignInputField(input, "TLSInspectionConfiguration", _networkfirewallTLSInspectionConfiguration); err != nil {
			log.Errorf("invalid --tls-inspection-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTLSInspectionConfigurationName) > 0 {
		input.TLSInspectionConfigurationName = aws.String(_networkfirewallTLSInspectionConfigurationName)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTLSInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a firewall endpoint for an Network Firewall firewall. This type of
// firewall endpoint is independent of the firewall endpoints that you specify in
// the Firewall itself, and you define it in addition to those endpoints after the
// firewall has been created. You can define a VPC endpoint association using a
// different VPC than the one you used in the firewall specifications.
func networkfirewall_CreateVpcEndpointAssociation(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.CreateVpcEndpointAssociationInput{
		// FirewallArn: *string, // Required
		// SubnetMapping: *types.SubnetMapping, // Required
		// VpcId: *string, // Required
	}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallSubnetMapping) > 0 {
		if err := assignInputField(input, "SubnetMapping", _networkfirewallSubnetMapping); err != nil {
			log.Errorf("invalid --subnet-mapping: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallVpcId) > 0 {
		input.VpcId = aws.String(_networkfirewallVpcId)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcEndpointAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Firewall and its FirewallStatus. This operation requires the firewall's
// DeleteProtection flag to be FALSE . You can't revert this operation.
//
// You can check whether a firewall is in use by reviewing the route tables for
// the Availability Zones where you have firewall subnet mappings. Retrieve the
// subnet mappings by calling DescribeFirewall. You define and update the route tables through
// Amazon VPC. As needed, update the route tables for the zones to remove the
// firewall endpoints. When the route tables no longer use the firewall endpoints,
// you can remove the firewall safely.
//
// To delete a firewall, remove the delete protection if you need to using UpdateFirewallDeleteProtection, then
// delete the firewall by calling DeleteFirewall.
func networkfirewall_DeleteFirewall(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteFirewallInput{}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}

	if resp, err := client.DeleteFirewall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified FirewallPolicy.
func networkfirewall_DeleteFirewallPolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteFirewallPolicyInput{}

	if len(_networkfirewallFirewallPolicyArn) > 0 {
		input.FirewallPolicyArn = aws.String(_networkfirewallFirewallPolicyArn)
	}
	if len(_networkfirewallFirewallPolicyName) > 0 {
		input.FirewallPolicyName = aws.String(_networkfirewallFirewallPolicyName)
	}

	if resp, err := client.DeleteFirewallPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a transit gateway attachment from a Network Firewall. Either the
// firewall owner or the transit gateway owner can delete the attachment.
//
// After you delete a transit gateway attachment, traffic will no longer flow
// through the firewall endpoints.
//
// After you initiate the delete operation, use DescribeFirewall to monitor the deletion status.
func networkfirewall_DeleteNetworkFirewallTransitGatewayAttachment(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteNetworkFirewallTransitGatewayAttachmentInput{
		// TransitGatewayAttachmentId: *string, // Required
	}

	if len(_networkfirewallTransitGatewayAttachmentId) > 0 {
		input.TransitGatewayAttachmentId = aws.String(_networkfirewallTransitGatewayAttachmentId)
	}

	if resp, err := client.DeleteNetworkFirewallTransitGatewayAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Proxy.
// Detaches a Proxy configuration from a NAT Gateway.
func networkfirewall_DeleteProxy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteProxyInput{
		// NatGatewayId: *string, // Required
	}

	if len(_networkfirewallNatGatewayId) > 0 {
		input.NatGatewayId = aws.String(_networkfirewallNatGatewayId)
	}
	if len(_networkfirewallProxyArn) > 0 {
		input.ProxyArn = aws.String(_networkfirewallProxyArn)
	}
	if len(_networkfirewallProxyName) > 0 {
		input.ProxyName = aws.String(_networkfirewallProxyName)
	}

	if resp, err := client.DeleteProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified ProxyConfiguration.
func networkfirewall_DeleteProxyConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteProxyConfigurationInput{}

	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}

	if resp, err := client.DeleteProxyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified ProxyRuleGroup.
func networkfirewall_DeleteProxyRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteProxyRuleGroupInput{}

	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}

	if resp, err := client.DeleteProxyRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified ProxyRule(s). currently attached to a ProxyRuleGroup
func networkfirewall_DeleteProxyRules(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteProxyRulesInput{
		// Rules: []string, // Required
	}

	if len(_networkfirewallRules) > 0 {
		input.Rules = []string{_networkfirewallRules}
	}
	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}

	if resp, err := client.DeleteProxyRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource policy that you created in a PutResourcePolicy request.
func networkfirewall_DeleteResourcePolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkfirewallResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkfirewallResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified RuleGroup.
func networkfirewall_DeleteRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteRuleGroupInput{}

	if len(_networkfirewallRuleGroupArn) > 0 {
		input.RuleGroupArn = aws.String(_networkfirewallRuleGroupArn)
	}
	if len(_networkfirewallRuleGroupName) > 0 {
		input.RuleGroupName = aws.String(_networkfirewallRuleGroupName)
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified TLSInspectionConfiguration.
func networkfirewall_DeleteTLSInspectionConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteTLSInspectionConfigurationInput{}

	if len(_networkfirewallTLSInspectionConfigurationArn) > 0 {
		input.TLSInspectionConfigurationArn = aws.String(_networkfirewallTLSInspectionConfigurationArn)
	}
	if len(_networkfirewallTLSInspectionConfigurationName) > 0 {
		input.TLSInspectionConfigurationName = aws.String(_networkfirewallTLSInspectionConfigurationName)
	}

	if resp, err := client.DeleteTLSInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified VpcEndpointAssociation.
// You can check whether an endpoint association is in use by reviewing the route
// tables for the Availability Zones where you have the endpoint subnet mapping.
// You can retrieve the subnet mapping by calling DescribeVpcEndpointAssociation. You define and update the
// route tables through Amazon VPC. As needed, update the route tables for the
// Availability Zone to remove the firewall endpoint for the association. When the
// route tables no longer use the firewall endpoint, you can remove the endpoint
// association safely.
func networkfirewall_DeleteVpcEndpointAssociation(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DeleteVpcEndpointAssociationInput{
		// VpcEndpointAssociationArn: *string, // Required
	}

	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}

	if resp, err := client.DeleteVpcEndpointAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified firewall.
func networkfirewall_DescribeFirewall(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeFirewallInput{}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}

	if resp, err := client.DescribeFirewall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the high-level information about a firewall, including the Availability
// Zones where the Firewall is currently in use.
func networkfirewall_DescribeFirewallMetadata(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeFirewallMetadataInput{}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}

	if resp, err := client.DescribeFirewallMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified firewall policy.
func networkfirewall_DescribeFirewallPolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeFirewallPolicyInput{}

	if len(_networkfirewallFirewallPolicyArn) > 0 {
		input.FirewallPolicyArn = aws.String(_networkfirewallFirewallPolicyArn)
	}
	if len(_networkfirewallFirewallPolicyName) > 0 {
		input.FirewallPolicyName = aws.String(_networkfirewallFirewallPolicyName)
	}

	if resp, err := client.DescribeFirewallPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns key information about a specific flow operation.
func networkfirewall_DescribeFlowOperation(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeFlowOperationInput{
		// FirewallArn: *string, // Required
		// FlowOperationId: *string, // Required
	}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFlowOperationId) > 0 {
		input.FlowOperationId = aws.String(_networkfirewallFlowOperationId)
	}
	if len(_networkfirewallAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_networkfirewallAvailabilityZone)
	}
	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}
	if len(_networkfirewallVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_networkfirewallVpcEndpointId)
	}

	if resp, err := client.DescribeFlowOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the logging configuration for the specified firewall.
func networkfirewall_DescribeLoggingConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeLoggingConfigurationInput{}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}

	if resp, err := client.DescribeLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified proxy.
func networkfirewall_DescribeProxy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeProxyInput{}

	if len(_networkfirewallProxyArn) > 0 {
		input.ProxyArn = aws.String(_networkfirewallProxyArn)
	}
	if len(_networkfirewallProxyName) > 0 {
		input.ProxyName = aws.String(_networkfirewallProxyName)
	}

	if resp, err := client.DescribeProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified proxy configuration.
func networkfirewall_DescribeProxyConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeProxyConfigurationInput{}

	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}

	if resp, err := client.DescribeProxyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified proxy configuration for the
// specified proxy rule group.
func networkfirewall_DescribeProxyRule(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeProxyRuleInput{
		// ProxyRuleName: *string, // Required
	}

	if len(_networkfirewallProxyRuleName) > 0 {
		input.ProxyRuleName = aws.String(_networkfirewallProxyRuleName)
	}
	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}

	if resp, err := client.DescribeProxyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified proxy rule group.
func networkfirewall_DescribeProxyRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeProxyRuleGroupInput{}

	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}

	if resp, err := client.DescribeProxyRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a resource policy that you created in a PutResourcePolicy request.
func networkfirewall_DescribeResourcePolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkfirewallResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkfirewallResourceArn)
	}

	if resp, err := client.DescribeResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified rule group.
func networkfirewall_DescribeRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeRuleGroupInput{}

	if len(_networkfirewallAnalyzeRuleGroup) > 0 {
		if err := assignInputField(input, "AnalyzeRuleGroup", _networkfirewallAnalyzeRuleGroup); err != nil {
			log.Errorf("invalid --analyze-rule-group: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRuleGroupArn) > 0 {
		input.RuleGroupArn = aws.String(_networkfirewallRuleGroupArn)
	}
	if len(_networkfirewallRuleGroupName) > 0 {
		input.RuleGroupName = aws.String(_networkfirewallRuleGroupName)
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// High-level information about a rule group, returned by operations like create
// and describe. You can use the information provided in the metadata to retrieve
// and manage a rule group. You can retrieve all objects for a rule group by
// calling DescribeRuleGroup.
func networkfirewall_DescribeRuleGroupMetadata(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeRuleGroupMetadataInput{}

	if len(_networkfirewallRuleGroupArn) > 0 {
		input.RuleGroupArn = aws.String(_networkfirewallRuleGroupArn)
	}
	if len(_networkfirewallRuleGroupName) > 0 {
		input.RuleGroupName = aws.String(_networkfirewallRuleGroupName)
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeRuleGroupMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information for a stateful rule group.
// For active threat defense Amazon Web Services managed rule groups, this
// operation provides insight into the protections enabled by the rule group, based
// on Suricata rule metadata fields. Summaries are available for rule groups you
// manage and for active threat defense Amazon Web Services managed rule groups.
//
// To modify how threat information appears in summaries, use the
// SummaryConfiguration parameter in UpdateRuleGroup.
func networkfirewall_DescribeRuleGroupSummary(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeRuleGroupSummaryInput{}

	if len(_networkfirewallRuleGroupArn) > 0 {
		input.RuleGroupArn = aws.String(_networkfirewallRuleGroupArn)
	}
	if len(_networkfirewallRuleGroupName) > 0 {
		input.RuleGroupName = aws.String(_networkfirewallRuleGroupName)
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeRuleGroupSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data objects for the specified TLS inspection configuration.
func networkfirewall_DescribeTLSInspectionConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeTLSInspectionConfigurationInput{}

	if len(_networkfirewallTLSInspectionConfigurationArn) > 0 {
		input.TLSInspectionConfigurationArn = aws.String(_networkfirewallTLSInspectionConfigurationArn)
	}
	if len(_networkfirewallTLSInspectionConfigurationName) > 0 {
		input.TLSInspectionConfigurationName = aws.String(_networkfirewallTLSInspectionConfigurationName)
	}

	if resp, err := client.DescribeTLSInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data object for the specified VPC endpoint association.
func networkfirewall_DescribeVpcEndpointAssociation(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DescribeVpcEndpointAssociationInput{
		// VpcEndpointAssociationArn: *string, // Required
	}

	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}

	if resp, err := client.DescribeVpcEndpointAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches ProxyRuleGroup resources from a ProxyConfiguration
// A Proxy Configuration defines the monitoring and protection behavior for a
// Proxy. The details of the behavior are defined in the rule groups that you add
// to your configuration.
func networkfirewall_DetachRuleGroupsFromProxyConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DetachRuleGroupsFromProxyConfigurationInput{
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}
	if len(_networkfirewallRuleGroupArns) > 0 {
		input.RuleGroupArns = append([]string(nil), _networkfirewallRuleGroupArns...)
	}
	if len(_networkfirewallRuleGroupNames) > 0 {
		input.RuleGroupNames = append([]string(nil), _networkfirewallRuleGroupNames...)
	}

	if resp, err := client.DetachRuleGroupsFromProxyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified Availability Zone associations from a transit
// gateway-attached firewall. This removes the firewall endpoints from these
// Availability Zones and stops traffic filtering in those zones. Before removing
// an Availability Zone, ensure you've updated your transit gateway route tables to
// redirect traffic appropriately.
//
// If AvailabilityZoneChangeProtection is enabled, you must first disable it using UpdateAvailabilityZoneChangeProtection
// .
//
// To verify the status of your Availability Zone changes, use DescribeFirewall.
func networkfirewall_DisassociateAvailabilityZones(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DisassociateAvailabilityZonesInput{
		// AvailabilityZoneMappings: []types.AvailabilityZoneMapping, // Required
	}

	if len(_networkfirewallAvailabilityZoneMappings) > 0 {
		if err := assignInputField(input, "AvailabilityZoneMappings", _networkfirewallAvailabilityZoneMappings); err != nil {
			log.Errorf("invalid --availability-zone-mappings: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.DisassociateAvailabilityZones(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified subnet associations from the firewall. This removes the
// firewall endpoints from the subnets and removes any network filtering
// protections that the endpoints were providing.
func networkfirewall_DisassociateSubnets(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.DisassociateSubnetsInput{
		// SubnetIds: []string, // Required
	}

	if len(_networkfirewallSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _networkfirewallSubnetIds...)
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.DisassociateSubnets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The results of a COMPLETED analysis report generated with StartAnalysisReport.
// For more information, see AnalysisTypeReportResult.
func networkfirewall_GetAnalysisReportResults(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.GetAnalysisReportResultsInput{
		// AnalysisReportId: *string, // Required
	}

	if len(_networkfirewallAnalysisReportId) > 0 {
		input.AnalysisReportId = aws.String(_networkfirewallAnalysisReportId)
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAnalysisReportResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.GetAnalysisReportResultsOutput
	p := networkfirewall.NewGetAnalysisReportResultsPaginator(client, input)
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

// Returns a list of all traffic analysis reports generated within the last 30
// days.
func networkfirewall_ListAnalysisReports(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListAnalysisReportsInput{}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnalysisReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListAnalysisReportsOutput
	p := networkfirewall.NewListAnalysisReportsPaginator(client, input)
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

// Retrieves the metadata for the firewall policies that you have defined.
// Depending on your setting for max results and the number of firewall policies, a
// single call might not return the full list.
func networkfirewall_ListFirewallPolicies(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListFirewallPoliciesInput{}

	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListFirewallPoliciesOutput
	p := networkfirewall.NewListFirewallPoliciesPaginator(client, input)
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

// Retrieves the metadata for the firewalls that you have defined. If you provide
// VPC identifiers in your request, this returns only the firewalls for those VPCs.
//
// Depending on your setting for max results and the number of firewalls, a single
// call might not return the full list.
func networkfirewall_ListFirewalls(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListFirewallsInput{}

	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}
	if len(_networkfirewallVpcIds) > 0 {
		input.VpcIds = append([]string(nil), _networkfirewallVpcIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewalls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListFirewallsOutput
	p := networkfirewall.NewListFirewallsPaginator(client, input)
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

// Returns the results of a specific flow operation.
// Flow operations let you manage the flows tracked in the flow table, also known
// as the firewall table.
//
// A flow is network traffic that is monitored by a firewall, either by stateful
// or stateless rules. For traffic to be considered part of a flow, it must share
// Destination, DestinationPort, Direction, Protocol, Source, and SourcePort.
func networkfirewall_ListFlowOperationResults(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListFlowOperationResultsInput{
		// FirewallArn: *string, // Required
		// FlowOperationId: *string, // Required
	}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFlowOperationId) > 0 {
		input.FlowOperationId = aws.String(_networkfirewallFlowOperationId)
	}
	if len(_networkfirewallAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_networkfirewallAvailabilityZone)
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}
	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}
	if len(_networkfirewallVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_networkfirewallVpcEndpointId)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowOperationResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListFlowOperationResultsOutput
	p := networkfirewall.NewListFlowOperationResultsPaginator(client, input)
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

// Returns a list of all flow operations ran in a specific firewall. You can
// optionally narrow the request scope by specifying the operation type or
// Availability Zone associated with a firewall's flow operations.
//
// Flow operations let you manage the flows tracked in the flow table, also known
// as the firewall table.
//
// A flow is network traffic that is monitored by a firewall, either by stateful
// or stateless rules. For traffic to be considered part of a flow, it must share
// Destination, DestinationPort, Direction, Protocol, Source, and SourcePort.
func networkfirewall_ListFlowOperations(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListFlowOperationsInput{
		// FirewallArn: *string, // Required
	}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_networkfirewallAvailabilityZone)
	}
	if len(_networkfirewallFlowOperationType) > 0 {
		if err := assignInputField(input, "FlowOperationType", _networkfirewallFlowOperationType); err != nil {
			log.Errorf("invalid --flow-operation-type: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}
	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}
	if len(_networkfirewallVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_networkfirewallVpcEndpointId)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListFlowOperationsOutput
	p := networkfirewall.NewListFlowOperationsPaginator(client, input)
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

// Retrieves the metadata for the proxies that you have defined. Depending on your
// setting for max results and the number of proxies, a single call might not
// return the full list.
func networkfirewall_ListProxies(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListProxiesInput{}

	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProxies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListProxiesOutput
	p := networkfirewall.NewListProxiesPaginator(client, input)
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

// Retrieves the metadata for the proxy configuration that you have defined.
// Depending on your setting for max results and the number of proxy
// configurations, a single call might not return the full list.
func networkfirewall_ListProxyConfigurations(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListProxyConfigurationsInput{}

	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProxyConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListProxyConfigurationsOutput
	p := networkfirewall.NewListProxyConfigurationsPaginator(client, input)
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

// Retrieves the metadata for the proxy rule groups that you have defined.
// Depending on your setting for max results and the number of proxy rule groups, a
// single call might not return the full list.
func networkfirewall_ListProxyRuleGroups(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListProxyRuleGroupsInput{}

	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProxyRuleGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListProxyRuleGroupsOutput
	p := networkfirewall.NewListProxyRuleGroupsPaginator(client, input)
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

// Retrieves the metadata for the rule groups that you have defined. Depending on
// your setting for max results and the number of rule groups, a single call might
// not return the full list.
func networkfirewall_ListRuleGroups(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListRuleGroupsInput{}

	if len(_networkfirewallManagedType) > 0 {
		if err := assignInputField(input, "ManagedType", _networkfirewallManagedType); err != nil {
			log.Errorf("invalid --managed-type: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}
	if len(_networkfirewallScope) > 0 {
		if err := assignInputField(input, "Scope", _networkfirewallScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallSubscriptionStatus) > 0 {
		if err := assignInputField(input, "SubscriptionStatus", _networkfirewallSubscriptionStatus); err != nil {
			log.Errorf("invalid --subscription-status: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRuleGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListRuleGroupsOutput
	p := networkfirewall.NewListRuleGroupsPaginator(client, input)
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

// Retrieves the tags associated with the specified resource. Tags are key:value
// pairs that you can use to categorize and manage your resources, for purposes
// like billing. For example, you might set the tag key to "customer" and the value
// to the customer name or ID. You can specify one or more tags to add to each
// Amazon Web Services resource, up to 50 tags for a resource.
//
// You can tag the Amazon Web Services resources that you manage through Network
// Firewall: firewalls, firewall policies, and rule groups.
func networkfirewall_ListTagsForResource(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkfirewallResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkfirewallResourceArn)
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
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

	var results []*networkfirewall.ListTagsForResourceOutput
	p := networkfirewall.NewListTagsForResourcePaginator(client, input)
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

// Retrieves the metadata for the TLS inspection configurations that you have
// defined. Depending on your setting for max results and the number of TLS
// inspection configurations, a single call might not return the full list.
func networkfirewall_ListTLSInspectionConfigurations(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListTLSInspectionConfigurationsInput{}

	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTLSInspectionConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListTLSInspectionConfigurationsOutput
	p := networkfirewall.NewListTLSInspectionConfigurationsPaginator(client, input)
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

// Retrieves the metadata for the VPC endpoint associations that you have defined.
// If you specify a fireawll, this returns only the endpoint associations for that
// firewall.
//
// Depending on your setting for max results and the number of associations, a
// single call might not return the full list.
func networkfirewall_ListVpcEndpointAssociations(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.ListVpcEndpointAssociationsInput{}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkfirewallMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallNextToken) > 0 {
		input.NextToken = aws.String(_networkfirewallNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVpcEndpointAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkfirewall.ListVpcEndpointAssociationsOutput
	p := networkfirewall.NewListVpcEndpointAssociationsPaginator(client, input)
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

// Creates or updates an IAM policy for your rule group, firewall policy, or
// firewall. Use this to share these resources between accounts. This operation
// works in conjunction with the Amazon Web Services Resource Access Manager (RAM)
// service to manage resource sharing for Network Firewall.
//
// For information about using sharing with Network Firewall resources, see [Sharing Network Firewall resources] in
// the Network Firewall Developer Guide.
//
// Use this operation to create or update a resource policy for your Network
// Firewall rule group, firewall policy, or firewall. In the resource policy, you
// specify the accounts that you want to share the Network Firewall resource with
// and the operations that you want the accounts to be able to perform.
//
// When you add an account in the resource policy, you then run the following
// Resource Access Manager (RAM) operations to access and accept the shared
// resource.
//
// [GetResourceShareInvitations]
// - - Returns the Amazon Resource Names (ARNs) of the resource share
// invitations.
//
// [AcceptResourceShareInvitation]
// - - Accepts the share invitation for a specified resource share.
//
// For additional information about resource sharing using RAM, see [Resource Access Manager User Guide].
//
// [AcceptResourceShareInvitation]: https://docs.aws.amazon.com/ram/latest/APIReference/API_AcceptResourceShareInvitation.html
// [GetResourceShareInvitations]: https://docs.aws.amazon.com/ram/latest/APIReference/API_GetResourceShareInvitations.html
// [Sharing Network Firewall resources]: https://docs.aws.amazon.com/network-firewall/latest/developerguide/sharing.html
// [Resource Access Manager User Guide]: https://docs.aws.amazon.com/ram/latest/userguide/what-is.html
func networkfirewall_PutResourcePolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_networkfirewallPolicy) > 0 {
		input.Policy = aws.String(_networkfirewallPolicy)
	}
	if len(_networkfirewallResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkfirewallResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a transit gateway attachment request for Network Firewall. When you
// reject the attachment request, Network Firewall cancels the creation of routing
// components between the transit gateway and firewall endpoints.
//
// Only the transit gateway owner can reject the attachment. After rejection, no
// traffic will flow through the firewall endpoints for this attachment.
//
// Use DescribeFirewall to monitor the rejection status. To accept the attachment instead of
// rejecting it, use AcceptNetworkFirewallTransitGatewayAttachment.
//
// Once rejected, you cannot reverse this action. To establish connectivity, you
// must create a new transit gateway-attached firewall.
func networkfirewall_RejectNetworkFirewallTransitGatewayAttachment(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.RejectNetworkFirewallTransitGatewayAttachmentInput{
		// TransitGatewayAttachmentId: *string, // Required
	}

	if len(_networkfirewallTransitGatewayAttachmentId) > 0 {
		input.TransitGatewayAttachmentId = aws.String(_networkfirewallTransitGatewayAttachmentId)
	}

	if resp, err := client.RejectNetworkFirewallTransitGatewayAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a traffic analysis report for the timeframe and traffic type you
// specify.
//
// For information on the contents of a traffic analysis report, see AnalysisReport.
func networkfirewall_StartAnalysisReport(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.StartAnalysisReportInput{
		// AnalysisType: types.EnabledAnalysisType, // Required
	}

	if len(_networkfirewallAnalysisType) > 0 {
		if err := assignInputField(input, "AnalysisType", _networkfirewallAnalysisType); err != nil {
			log.Errorf("invalid --analysis-type: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}

	if resp, err := client.StartAnalysisReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins capturing the flows in a firewall, according to the filters you define.
// Captures are similar, but not identical to snapshots. Capture operations provide
// visibility into flows that are not closed and are tracked by a firewall's flow
// table. Unlike snapshots, captures are a time-boxed view.
//
// A flow is network traffic that is monitored by a firewall, either by stateful
// or stateless rules. For traffic to be considered part of a flow, it must share
// Destination, DestinationPort, Direction, Protocol, Source, and SourcePort.
//
// To avoid encountering operation limits, you should avoid starting captures with
// broad filters, like wide IP ranges. Instead, we recommend you define more
// specific criteria with FlowFilters , like narrow IP ranges, ports, or protocols.
func networkfirewall_StartFlowCapture(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.StartFlowCaptureInput{
		// FirewallArn: *string, // Required
		// FlowFilters: []types.FlowFilter, // Required
	}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFlowFilters) > 0 {
		if err := assignInputField(input, "FlowFilters", _networkfirewallFlowFilters); err != nil {
			log.Errorf("invalid --flow-filters: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_networkfirewallAvailabilityZone)
	}
	if len(_networkfirewallMinimumFlowAgeInSeconds) > 0 {
		if err := assignInputField(input, "MinimumFlowAgeInSeconds", _networkfirewallMinimumFlowAgeInSeconds); err != nil {
			log.Errorf("invalid --minimum-flow-age-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}
	if len(_networkfirewallVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_networkfirewallVpcEndpointId)
	}

	if resp, err := client.StartFlowCapture(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins the flushing of traffic from the firewall, according to the filters you
// define. When the operation starts, impacted flows are temporarily marked as
// timed out before the Suricata engine prunes, or flushes, the flows from the
// firewall table.
//
// While the flush completes, impacted flows are processed as midstream traffic.
// This may result in a temporary increase in midstream traffic metrics. We
// recommend that you double check your stream exception policy before you perform
// a flush operation.
func networkfirewall_StartFlowFlush(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.StartFlowFlushInput{
		// FirewallArn: *string, // Required
		// FlowFilters: []types.FlowFilter, // Required
	}

	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFlowFilters) > 0 {
		if err := assignInputField(input, "FlowFilters", _networkfirewallFlowFilters); err != nil {
			log.Errorf("invalid --flow-filters: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_networkfirewallAvailabilityZone)
	}
	if len(_networkfirewallMinimumFlowAgeInSeconds) > 0 {
		if err := assignInputField(input, "MinimumFlowAgeInSeconds", _networkfirewallMinimumFlowAgeInSeconds); err != nil {
			log.Errorf("invalid --minimum-flow-age-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallVpcEndpointAssociationArn) > 0 {
		input.VpcEndpointAssociationArn = aws.String(_networkfirewallVpcEndpointAssociationArn)
	}
	if len(_networkfirewallVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_networkfirewallVpcEndpointId)
	}

	if resp, err := client.StartFlowFlush(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource. Tags are key:value pairs
// that you can use to categorize and manage your resources, for purposes like
// billing. For example, you might set the tag key to "customer" and the value to
// the customer name or ID. You can specify one or more tags to add to each Amazon
// Web Services resource, up to 50 tags for a resource.
//
// You can tag the Amazon Web Services resources that you manage through Network
// Firewall: firewalls, firewall policies, and rule groups.
func networkfirewall_TagResource(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_networkfirewallResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkfirewallResourceArn)
	}
	if len(_networkfirewallTags) > 0 {
		if err := assignInputField(input, "Tags", _networkfirewallTags); err != nil {
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

// Removes the tags with the specified keys from the specified resource. Tags are
// key:value pairs that you can use to categorize and manage your resources, for
// purposes like billing. For example, you might set the tag key to "customer" and
// the value to the customer name or ID. You can specify one or more tags to add to
// each Amazon Web Services resource, up to 50 tags for a resource.
//
// You can manage tags for the Amazon Web Services resources that you manage
// through Network Firewall: firewalls, firewall policies, and rule groups.
func networkfirewall_UntagResource(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_networkfirewallResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkfirewallResourceArn)
	}
	if len(_networkfirewallTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _networkfirewallTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the AvailabilityZoneChangeProtection setting for a transit
// gateway-attached firewall. When enabled, this setting prevents accidental
// changes to the firewall's Availability Zone configuration. This helps protect
// against disrupting traffic flow in production environments.
//
// When enabled, you must disable this protection before using AssociateAvailabilityZones or DisassociateAvailabilityZones to modify the
// firewall's Availability Zone configuration.
func networkfirewall_UpdateAvailabilityZoneChangeProtection(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateAvailabilityZoneChangeProtectionInput{
		// AvailabilityZoneChangeProtection: bool, // Required
	}

	if len(_networkfirewallAvailabilityZoneChangeProtection) > 0 {
		if err := assignInputField(input, "AvailabilityZoneChangeProtection", _networkfirewallAvailabilityZoneChangeProtection); err != nil {
			log.Errorf("invalid --availability-zone-change-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateAvailabilityZoneChangeProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables specific types of firewall analysis on a specific firewall you define.
func networkfirewall_UpdateFirewallAnalysisSettings(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateFirewallAnalysisSettingsInput{}

	if len(_networkfirewallEnabledAnalysisTypes) > 0 {
		if err := assignInputField(input, "EnabledAnalysisTypes", _networkfirewallEnabledAnalysisTypes); err != nil {
			log.Errorf("invalid --enabled-analysis-types: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateFirewallAnalysisSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the flag, DeleteProtection , which indicates whether it is possible to
// delete the firewall. If the flag is set to TRUE , the firewall is protected
// against deletion. This setting helps protect against accidentally deleting a
// firewall that's in use.
func networkfirewall_UpdateFirewallDeleteProtection(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateFirewallDeleteProtectionInput{
		// DeleteProtection: bool, // Required
	}

	if len(_networkfirewallDeleteProtection) > 0 {
		if err := assignInputField(input, "DeleteProtection", _networkfirewallDeleteProtection); err != nil {
			log.Errorf("invalid --delete-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateFirewallDeleteProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the description for the specified firewall. Use the description to
// help you identify the firewall when you're working with it.
func networkfirewall_UpdateFirewallDescription(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateFirewallDescriptionInput{}

	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateFirewallDescription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A complex type that contains settings for encryption of your firewall resources.
func networkfirewall_UpdateFirewallEncryptionConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateFirewallEncryptionConfigurationInput{}

	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateFirewallEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of the specified firewall policy.
func networkfirewall_UpdateFirewallPolicy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateFirewallPolicyInput{
		// FirewallPolicy: *types.FirewallPolicy, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallFirewallPolicy) > 0 {
		if err := assignInputField(input, "FirewallPolicy", _networkfirewallFirewallPolicy); err != nil {
			log.Errorf("invalid --firewall-policy: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _networkfirewallDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallPolicyArn) > 0 {
		input.FirewallPolicyArn = aws.String(_networkfirewallFirewallPolicyArn)
	}
	if len(_networkfirewallFirewallPolicyName) > 0 {
		input.FirewallPolicyName = aws.String(_networkfirewallFirewallPolicyName)
	}

	if resp, err := client.UpdateFirewallPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the flag, ChangeProtection , which indicates whether it is possible to
// change the firewall. If the flag is set to TRUE , the firewall is protected from
// changes. This setting helps protect against accidentally changing a firewall
// that's in use.
func networkfirewall_UpdateFirewallPolicyChangeProtection(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateFirewallPolicyChangeProtectionInput{
		// FirewallPolicyChangeProtection: bool, // Required
	}

	if len(_networkfirewallFirewallPolicyChangeProtection) > 0 {
		if err := assignInputField(input, "FirewallPolicyChangeProtection", _networkfirewallFirewallPolicyChangeProtection); err != nil {
			log.Errorf("invalid --firewall-policy-change-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateFirewallPolicyChangeProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the logging configuration for the specified firewall.
// To change the logging configuration, retrieve the LoggingConfiguration by calling DescribeLoggingConfiguration, then change it
// and provide the modified object to this update call. You must change the logging
// configuration one LogDestinationConfigat a time inside the retrieved LoggingConfiguration object.
//
// You can perform only one of the following actions in any call to
// UpdateLoggingConfiguration :
//
// - Create a new log destination object by adding a single LogDestinationConfig
// array element to LogDestinationConfigs .
//
// - Delete a log destination object by removing a single LogDestinationConfig
// array element from LogDestinationConfigs .
//
// - Change the LogDestination setting in a single LogDestinationConfig array
// element.
//
// You can't change the LogDestinationType or LogType in a LogDestinationConfig .
// To change these settings, delete the existing LogDestinationConfig object and
// create a new one, using two separate calls to this update operation.
func networkfirewall_UpdateLoggingConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateLoggingConfigurationInput{}

	if len(_networkfirewallEnableMonitoringDashboard) > 0 {
		if err := assignInputField(input, "EnableMonitoringDashboard", _networkfirewallEnableMonitoringDashboard); err != nil {
			log.Errorf("invalid --enable-monitoring-dashboard: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _networkfirewallLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of the specified proxy.
func networkfirewall_UpdateProxy(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateProxyInput{
		// NatGatewayId: *string, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallNatGatewayId) > 0 {
		input.NatGatewayId = aws.String(_networkfirewallNatGatewayId)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallListenerPropertiesToAdd) > 0 {
		if err := assignInputField(input, "ListenerPropertiesToAdd", _networkfirewallListenerPropertiesToAdd); err != nil {
			log.Errorf("invalid --listener-properties-to-add: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallListenerPropertiesToRemove) > 0 {
		if err := assignInputField(input, "ListenerPropertiesToRemove", _networkfirewallListenerPropertiesToRemove); err != nil {
			log.Errorf("invalid --listener-properties-to-remove: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallProxyArn) > 0 {
		input.ProxyArn = aws.String(_networkfirewallProxyArn)
	}
	if len(_networkfirewallProxyName) > 0 {
		input.ProxyName = aws.String(_networkfirewallProxyName)
	}
	if len(_networkfirewallTlsInterceptProperties) > 0 {
		if err := assignInputField(input, "TlsInterceptProperties", _networkfirewallTlsInterceptProperties); err != nil {
			log.Errorf("invalid --tls-intercept-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of the specified proxy configuration.
func networkfirewall_UpdateProxyConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateProxyConfigurationInput{
		// DefaultRulePhaseActions: *types.ProxyConfigDefaultRulePhaseActionsRequest, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallDefaultRulePhaseActions) > 0 {
		if err := assignInputField(input, "DefaultRulePhaseActions", _networkfirewallDefaultRulePhaseActions); err != nil {
			log.Errorf("invalid --default-rule-phase-actions: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}

	if resp, err := client.UpdateProxyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of the specified proxy rule.
func networkfirewall_UpdateProxyRule(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateProxyRuleInput{
		// ProxyRuleName: *string, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallProxyRuleName) > 0 {
		input.ProxyRuleName = aws.String(_networkfirewallProxyRuleName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallAction) > 0 {
		if err := assignInputField(input, "Action", _networkfirewallAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallAddConditions) > 0 {
		if err := assignInputField(input, "AddConditions", _networkfirewallAddConditions); err != nil {
			log.Errorf("invalid --add-conditions: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}
	if len(_networkfirewallRemoveConditions) > 0 {
		if err := assignInputField(input, "RemoveConditions", _networkfirewallRemoveConditions); err != nil {
			log.Errorf("invalid --remove-conditions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProxyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates proxy rule group priorities within a proxy configuration.
func networkfirewall_UpdateProxyRuleGroupPriorities(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateProxyRuleGroupPrioritiesInput{
		// RuleGroups: []types.ProxyRuleGroupPriority, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallRuleGroups) > 0 {
		if err := assignInputField(input, "RuleGroups", _networkfirewallRuleGroups); err != nil {
			log.Errorf("invalid --rule-groups: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallProxyConfigurationArn) > 0 {
		input.ProxyConfigurationArn = aws.String(_networkfirewallProxyConfigurationArn)
	}
	if len(_networkfirewallProxyConfigurationName) > 0 {
		input.ProxyConfigurationName = aws.String(_networkfirewallProxyConfigurationName)
	}

	if resp, err := client.UpdateProxyRuleGroupPriorities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates proxy rule priorities within a proxy rule group.
func networkfirewall_UpdateProxyRulePriorities(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateProxyRulePrioritiesInput{
		// RuleGroupRequestPhase: types.RuleGroupRequestPhase, // Required
		// Rules: []types.ProxyRulePriority, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallRuleGroupRequestPhase) > 0 {
		if err := assignInputField(input, "RuleGroupRequestPhase", _networkfirewallRuleGroupRequestPhase); err != nil {
			log.Errorf("invalid --rule-group-request-phase: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRules) > 0 {
		if err := assignInputField(input, "Rules", _networkfirewallRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallProxyRuleGroupArn) > 0 {
		input.ProxyRuleGroupArn = aws.String(_networkfirewallProxyRuleGroupArn)
	}
	if len(_networkfirewallProxyRuleGroupName) > 0 {
		input.ProxyRuleGroupName = aws.String(_networkfirewallProxyRuleGroupName)
	}

	if resp, err := client.UpdateProxyRulePriorities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the rule settings for the specified rule group. You use a rule group by
// reference in one or more firewall policies. When you modify a rule group, you
// modify all firewall policies that use the rule group.
//
// To update a rule group, first call DescribeRuleGroup to retrieve the current RuleGroup object, update the
// object as needed, and then provide the updated object to this call.
func networkfirewall_UpdateRuleGroup(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateRuleGroupInput{
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallAnalyzeRuleGroup) > 0 {
		if err := assignInputField(input, "AnalyzeRuleGroup", _networkfirewallAnalyzeRuleGroup); err != nil {
			log.Errorf("invalid --analyze-rule-group: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _networkfirewallDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRuleGroup) > 0 {
		if err := assignInputField(input, "RuleGroup", _networkfirewallRuleGroup); err != nil {
			log.Errorf("invalid --rule-group: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallRuleGroupArn) > 0 {
		input.RuleGroupArn = aws.String(_networkfirewallRuleGroupArn)
	}
	if len(_networkfirewallRuleGroupName) > 0 {
		input.RuleGroupName = aws.String(_networkfirewallRuleGroupName)
	}
	if len(_networkfirewallRules) > 0 {
		input.Rules = aws.String(_networkfirewallRules)
	}
	if len(_networkfirewallSourceMetadata) > 0 {
		if err := assignInputField(input, "SourceMetadata", _networkfirewallSourceMetadata); err != nil {
			log.Errorf("invalid --source-metadata: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallSummaryConfiguration) > 0 {
		if err := assignInputField(input, "SummaryConfiguration", _networkfirewallSummaryConfiguration); err != nil {
			log.Errorf("invalid --summary-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallType) > 0 {
		if err := assignInputField(input, "Type", _networkfirewallType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func networkfirewall_UpdateSubnetChangeProtection(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateSubnetChangeProtectionInput{
		// SubnetChangeProtection: bool, // Required
	}

	if len(_networkfirewallSubnetChangeProtection) > 0 {
		if err := assignInputField(input, "SubnetChangeProtection", _networkfirewallSubnetChangeProtection); err != nil {
			log.Errorf("invalid --subnet-change-protection: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallFirewallArn) > 0 {
		input.FirewallArn = aws.String(_networkfirewallFirewallArn)
	}
	if len(_networkfirewallFirewallName) > 0 {
		input.FirewallName = aws.String(_networkfirewallFirewallName)
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}

	if resp, err := client.UpdateSubnetChangeProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the TLS inspection configuration settings for the specified TLS
// inspection configuration. You use a TLS inspection configuration by referencing
// it in one or more firewall policies. When you modify a TLS inspection
// configuration, you modify all firewall policies that use the TLS inspection
// configuration.
//
// To update a TLS inspection configuration, first call DescribeTLSInspectionConfiguration to retrieve the current TLSInspectionConfiguration
// object, update the object as needed, and then provide the updated object to this
// call.
func networkfirewall_UpdateTLSInspectionConfiguration(cfg aws.Config, client *networkfirewall.Client) {
	input := &networkfirewall.UpdateTLSInspectionConfigurationInput{
		// TLSInspectionConfiguration: *types.TLSInspectionConfiguration, // Required
		// UpdateToken: *string, // Required
	}

	if len(_networkfirewallTLSInspectionConfiguration) > 0 {
		if err := assignInputField(input, "TLSInspectionConfiguration", _networkfirewallTLSInspectionConfiguration); err != nil {
			log.Errorf("invalid --tls-inspection-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallUpdateToken) > 0 {
		input.UpdateToken = aws.String(_networkfirewallUpdateToken)
	}
	if len(_networkfirewallDescription) > 0 {
		input.Description = aws.String(_networkfirewallDescription)
	}
	if len(_networkfirewallEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _networkfirewallEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_networkfirewallTLSInspectionConfigurationArn) > 0 {
		input.TLSInspectionConfigurationArn = aws.String(_networkfirewallTLSInspectionConfigurationArn)
	}
	if len(_networkfirewallTLSInspectionConfigurationName) > 0 {
		input.TLSInspectionConfigurationName = aws.String(_networkfirewallTLSInspectionConfigurationName)
	}

	if resp, err := client.UpdateTLSInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_networkfirewallCmd)
	_networkfirewallCmd.Flags().SortFlags = false

	_networkfirewallCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_networkfirewallCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_networkfirewallCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAction, "action", "", "", "Action")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAddConditions, "add-conditions", "", "", "Add Conditions")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAnalysisReportId, "analysis-report-id", "", "", "Analysis Report ID")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAnalysisType, "analysis-type", "", "", "Analysis Type")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAnalyzeRuleGroup, "analyze-rule-group", "", "", "Analyze Rule Group")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAvailabilityZoneChangeProtection, "availability-zone-change-protection", "", "", "Availability Zone Change Protection")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallAvailabilityZoneMappings, "availability-zone-mappings", "", "", "Availability Zone Mappings")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallCapacity, "capacity", "", "", "Capacity")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallDefaultRulePhaseActions, "default-rule-phase-actions", "", "", "Default Rule Phase Actions")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallDeleteProtection, "delete-protection", "", "", "Delete Protection")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallDescription, "description", "", "", "Description")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallDryRun, "dry-run", "", "", "Dry Run")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallEnableMonitoringDashboard, "enable-monitoring-dashboard", "", "", "Enable Monitoring Dashboard")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallEnabledAnalysisTypes, "enabled-analysis-types", "", "", "Enabled Analysis Types")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFirewallArn, "firewall-arn", "", "", "Firewall ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFirewallName, "firewall-name", "", "", "Firewall Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFirewallPolicy, "firewall-policy", "", "", "Firewall Policy")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFirewallPolicyArn, "firewall-policy-arn", "", "", "Firewall Policy ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFirewallPolicyChangeProtection, "firewall-policy-change-protection", "", "", "Firewall Policy Change Protection")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFirewallPolicyName, "firewall-policy-name", "", "", "Firewall Policy Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFlowFilters, "flow-filters", "", "", "Flow Filters")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFlowOperationId, "flow-operation-id", "", "", "Flow Operation ID")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallFlowOperationType, "flow-operation-type", "", "", "Flow Operation Type")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallListenerProperties, "listener-properties", "", "", "Listener Properties")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallListenerPropertiesToAdd, "listener-properties-to-add", "", "", "Listener Properties To Add")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallListenerPropertiesToRemove, "listener-properties-to-remove", "", "", "Listener Properties To Remove")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallManagedType, "managed-type", "", "", "Managed Type")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallMaxResults, "max-results", "", "", "Max Results")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallMinimumFlowAgeInSeconds, "minimum-flow-age-in-seconds", "", "", "Minimum Flow Age In Seconds")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallNatGatewayId, "nat-gateway-id", "", "", "NAT Gateway ID")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallNextToken, "next-token", "", "", "Next Token")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallPolicy, "policy", "", "", "Policy")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyArn, "proxy-arn", "", "", "Proxy ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyConfigurationArn, "proxy-configuration-arn", "", "", "Proxy Configuration ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyConfigurationName, "proxy-configuration-name", "", "", "Proxy Configuration Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyName, "proxy-name", "", "", "Proxy Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyRuleGroupArn, "proxy-rule-group-arn", "", "", "Proxy Rule Group ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyRuleGroupName, "proxy-rule-group-name", "", "", "Proxy Rule Group Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallProxyRuleName, "proxy-rule-name", "", "", "Proxy Rule Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRemoveConditions, "remove-conditions", "", "", "Remove Conditions")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallResourceArn, "resource-arn", "", "", "Resource ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRuleGroup, "rule-group", "", "", "Rule Group")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRuleGroupArn, "rule-group-arn", "", "", "Rule Group ARN")
	_networkfirewallCmd.Flags().StringSliceVarP(&_networkfirewallRuleGroupArns, "rule-group-arns", "", nil, "Rule Group Arns")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRuleGroupName, "rule-group-name", "", "", "Rule Group Name")
	_networkfirewallCmd.Flags().StringSliceVarP(&_networkfirewallRuleGroupNames, "rule-group-names", "", nil, "Rule Group Names")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRuleGroupRequestPhase, "rule-group-request-phase", "", "", "Rule Group Request Phase")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRuleGroups, "rule-groups", "", "", "Rule Groups")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallRules, "rules", "", "", "Rules")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallScope, "scope", "", "", "Scope")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallSourceMetadata, "source-metadata", "", "", "Source Metadata")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallSubnetChangeProtection, "subnet-change-protection", "", "", "Subnet Change Protection")
	_networkfirewallCmd.Flags().StringSliceVarP(&_networkfirewallSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallSubnetMapping, "subnet-mapping", "", "", "Subnet Mapping")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallSubnetMappings, "subnet-mappings", "", "", "Subnet Mappings")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallSubscriptionStatus, "subscription-status", "", "", "Subscription Status")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallSummaryConfiguration, "summary-configuration", "", "", "Summary Configuration")
	_networkfirewallCmd.Flags().StringSliceVarP(&_networkfirewallTagKeys, "tag-keys", "", nil, "Tag Keys")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTags, "tags", "", "", "Tags")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTLSInspectionConfiguration, "tls-inspection-configuration", "", "", "TLS Inspection Configuration")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTLSInspectionConfigurationArn, "tls-inspection-configuration-arn", "", "", "TLS Inspection Configuration ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTLSInspectionConfigurationName, "tls-inspection-configuration-name", "", "", "TLS Inspection Configuration Name")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTlsInterceptProperties, "tls-intercept-properties", "", "", "TLS Intercept Properties")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTransitGatewayAttachmentId, "transit-gateway-attachment-id", "", "", "Transit Gateway Attachment ID")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallTransitGatewayId, "transit-gateway-id", "", "", "Transit Gateway ID")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallType, "type", "", "", "Type")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallUpdateToken, "update-token", "", "", "Update Token")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallVpcEndpointAssociationArn, "vpc-endpoint-association-arn", "", "", "VPC Endpoint Association ARN")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallVpcEndpointId, "vpc-endpoint-id", "", "", "VPC Endpoint ID")
	_networkfirewallCmd.Flags().StringVarP(&_networkfirewallVpcId, "vpc-id", "", "", "VPC ID")
	_networkfirewallCmd.Flags().StringSliceVarP(&_networkfirewallVpcIds, "vpc-ids", "", nil, "VPC Ids")

	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallAcceptNetworkFirewallTransitGatewayAttachment, "accept-network-firewall-transit-gateway-attachment", "", false, "Accept Network Firewall Transit Gateway Attachment")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallAssociateAvailabilityZones, "associate-availability-zones", "", false, "Associate Availability Zones")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallAssociateFirewallPolicy, "associate-firewall-policy", "", false, "Associate Firewall Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallAssociateSubnets, "associate-subnets", "", false, "Associate Subnets")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallAttachRuleGroupsToProxyConfiguration, "attach-rule-groups-to-proxy-configuration", "", false, "Attach Rule Groups To Proxy Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateFirewall, "create-firewall", "", false, "Create Firewall")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateFirewallPolicy, "create-firewall-policy", "", false, "Create Firewall Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateProxy, "create-proxy", "", false, "Create Proxy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateProxyConfiguration, "create-proxy-configuration", "", false, "Create Proxy Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateProxyRuleGroup, "create-proxy-rule-group", "", false, "Create Proxy Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateProxyRules, "create-proxy-rules", "", false, "Create Proxy Rules")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateRuleGroup, "create-rule-group", "", false, "Create Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateTLSInspectionConfiguration, "create-tls-inspection-configuration", "", false, "Create TLS Inspection Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallCreateVpcEndpointAssociation, "create-vpc-endpoint-association", "", false, "Create VPC Endpoint Association")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteFirewall, "delete-firewall", "", false, "Delete Firewall")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteFirewallPolicy, "delete-firewall-policy", "", false, "Delete Firewall Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteNetworkFirewallTransitGatewayAttachment, "delete-network-firewall-transit-gateway-attachment", "", false, "Delete Network Firewall Transit Gateway Attachment")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteProxy, "delete-proxy", "", false, "Delete Proxy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteProxyConfiguration, "delete-proxy-configuration", "", false, "Delete Proxy Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteProxyRuleGroup, "delete-proxy-rule-group", "", false, "Delete Proxy Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteProxyRules, "delete-proxy-rules", "", false, "Delete Proxy Rules")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteRuleGroup, "delete-rule-group", "", false, "Delete Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteTLSInspectionConfiguration, "delete-tls-inspection-configuration", "", false, "Delete TLS Inspection Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDeleteVpcEndpointAssociation, "delete-vpc-endpoint-association", "", false, "Delete VPC Endpoint Association")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeFirewall, "describe-firewall", "", false, "Describe Firewall")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeFirewallMetadata, "describe-firewall-metadata", "", false, "Describe Firewall Metadata")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeFirewallPolicy, "describe-firewall-policy", "", false, "Describe Firewall Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeFlowOperation, "describe-flow-operation", "", false, "Describe Flow Operation")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeLoggingConfiguration, "describe-logging-configuration", "", false, "Describe Logging Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeProxy, "describe-proxy", "", false, "Describe Proxy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeProxyConfiguration, "describe-proxy-configuration", "", false, "Describe Proxy Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeProxyRule, "describe-proxy-rule", "", false, "Describe Proxy Rule")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeProxyRuleGroup, "describe-proxy-rule-group", "", false, "Describe Proxy Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeResourcePolicy, "describe-resource-policy", "", false, "Describe Resource Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeRuleGroup, "describe-rule-group", "", false, "Describe Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeRuleGroupMetadata, "describe-rule-group-metadata", "", false, "Describe Rule Group Metadata")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeRuleGroupSummary, "describe-rule-group-summary", "", false, "Describe Rule Group Summary")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeTLSInspectionConfiguration, "describe-tls-inspection-configuration", "", false, "Describe TLS Inspection Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDescribeVpcEndpointAssociation, "describe-vpc-endpoint-association", "", false, "Describe VPC Endpoint Association")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDetachRuleGroupsFromProxyConfiguration, "detach-rule-groups-from-proxy-configuration", "", false, "Detach Rule Groups From Proxy Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDisassociateAvailabilityZones, "disassociate-availability-zones", "", false, "Disassociate Availability Zones")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallDisassociateSubnets, "disassociate-subnets", "", false, "Disassociate Subnets")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallGetAnalysisReportResults, "get-analysis-report-results", "", false, "Get Analysis Report Results")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListAnalysisReports, "list-analysis-reports", "", false, "List Analysis Reports")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListFirewallPolicies, "list-firewall-policies", "", false, "List Firewall Policies")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListFirewalls, "list-firewalls", "", false, "List Firewalls")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListFlowOperationResults, "list-flow-operation-results", "", false, "List Flow Operation Results")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListFlowOperations, "list-flow-operations", "", false, "List Flow Operations")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListProxies, "list-proxies", "", false, "List Proxies")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListProxyConfigurations, "list-proxy-configurations", "", false, "List Proxy Configurations")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListProxyRuleGroups, "list-proxy-rule-groups", "", false, "List Proxy Rule Groups")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListRuleGroups, "list-rule-groups", "", false, "List Rule Groups")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListTLSInspectionConfigurations, "list-tls-inspection-configurations", "", false, "List TLS Inspection Configurations")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallListVpcEndpointAssociations, "list-vpc-endpoint-associations", "", false, "List VPC Endpoint Associations")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallRejectNetworkFirewallTransitGatewayAttachment, "reject-network-firewall-transit-gateway-attachment", "", false, "Reject Network Firewall Transit Gateway Attachment")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallStartAnalysisReport, "start-analysis-report", "", false, "Start Analysis Report")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallStartFlowCapture, "start-flow-capture", "", false, "Start Flow Capture")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallStartFlowFlush, "start-flow-flush", "", false, "Start Flow Flush")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallTagResource, "tag-resource", "", false, "Tag Resource")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUntagResource, "untag-resource", "", false, "Untag Resource")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateAvailabilityZoneChangeProtection, "update-availability-zone-change-protection", "", false, "Update Availability Zone Change Protection")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateFirewallAnalysisSettings, "update-firewall-analysis-settings", "", false, "Update Firewall Analysis Settings")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateFirewallDeleteProtection, "update-firewall-delete-protection", "", false, "Update Firewall Delete Protection")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateFirewallDescription, "update-firewall-description", "", false, "Update Firewall Description")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateFirewallEncryptionConfiguration, "update-firewall-encryption-configuration", "", false, "Update Firewall Encryption Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateFirewallPolicy, "update-firewall-policy", "", false, "Update Firewall Policy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateFirewallPolicyChangeProtection, "update-firewall-policy-change-protection", "", false, "Update Firewall Policy Change Protection")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateLoggingConfiguration, "update-logging-configuration", "", false, "Update Logging Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateProxy, "update-proxy", "", false, "Update Proxy")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateProxyConfiguration, "update-proxy-configuration", "", false, "Update Proxy Configuration")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateProxyRule, "update-proxy-rule", "", false, "Update Proxy Rule")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateProxyRuleGroupPriorities, "update-proxy-rule-group-priorities", "", false, "Update Proxy Rule Group Priorities")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateProxyRulePriorities, "update-proxy-rule-priorities", "", false, "Update Proxy Rule Priorities")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateRuleGroup, "update-rule-group", "", false, "Update Rule Group")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateSubnetChangeProtection, "update-subnet-change-protection", "", false, "Update Subnet Change Protection")
	_networkfirewallCmd.Flags().BoolVarP(&_networkfirewallUpdateTLSInspectionConfiguration, "update-tls-inspection-configuration", "", false, "Update TLS Inspection Configuration")

}
