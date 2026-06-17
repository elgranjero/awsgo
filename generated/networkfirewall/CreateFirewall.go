package networkfirewall

// CreateFirewall is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Creates an Network Firewall Firewall and accompanying FirewallStatus for a VPC.
//
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
