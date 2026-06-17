package route53resolver

// ImportFirewallDomains is generated as a reference stub.
// Executable command wiring lives under cmd/route53resolver.go.
//
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
