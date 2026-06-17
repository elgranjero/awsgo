package networkfirewall

// CreateTLSInspectionConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
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
