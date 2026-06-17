package fms

// PutPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/fms.go.
//
// Creates an Firewall Manager policy.
//
// A Firewall Manager policy is specific to the individual policy type. If you
// want to enforce multiple policy types across accounts, you can create multiple
// policies. You can create more than one policy for each type.
//
// If you add a new account to an organization that you created with
// Organizations, Firewall Manager automatically applies the policy to the
// resources in that account that are within scope of the policy.
//
// Firewall Manager provides the following types of policies:
//
// - WAF policy - This policy applies WAF web ACL protections to specified
// accounts and resources.
//
// - Shield Advanced policy - This policy applies Shield Advanced protection to
// specified accounts and resources.
//
// - Security Groups policy - This type of policy gives you control over
// security groups that are in use throughout your organization in Organizations
// and lets you enforce a baseline set of rules across your organization.
//
// - Network ACL policy - This type of policy gives you control over the network
// ACLs that are in use throughout your organization in Organizations and lets you
// enforce a baseline set of first and last network ACL rules across your
// organization.
//
// - Network Firewall policy - This policy applies Network Firewall protection
// to your organization's VPCs.
//
// - DNS Firewall policy - This policy applies Amazon Route 53 Resolver DNS
// Firewall protections to your organization's VPCs.
//
// - Third-party firewall policy - This policy applies third-party firewall
// protections. Third-party firewalls are available by subscription through the
// Amazon Web Services Marketplace console at [Amazon Web Services Marketplace].
//
// - Palo Alto Networks Cloud NGFW policy - This policy applies Palo Alto
// Networks Cloud Next Generation Firewall (NGFW) protections and Palo Alto
// Networks Cloud NGFW rulestacks to your organization's VPCs.
//
// - Fortigate CNF policy - This policy applies Fortigate Cloud Native Firewall
// (CNF) protections. Fortigate CNF is a cloud-centered solution that blocks
// Zero-Day threats and secures cloud infrastructures with industry-leading
// advanced threat prevention, smart web application firewalls (WAF), and API
// protection.
//
// [Amazon Web Services Marketplace]: http://aws.amazon.com/marketplace
