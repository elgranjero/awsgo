package organizations

// DetachPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Detaches a policy from a target root, organizational unit (OU), or account.
//
// If the policy being detached is a service control policy (SCP), the changes to
// permissions for Identity and Access Management (IAM) users and roles in affected
// accounts are immediate.
//
// Every root, OU, and account must have at least one SCP attached. If you want to
// replace the default FullAWSAccess policy with an SCP that limits the
// permissions that can be delegated, you must attach the replacement SCP before
// you can remove the default SCP. This is the authorization strategy of an "[allow list] ". If
// you instead attach a second SCP and leave the FullAWSAccess SCP still attached,
// and specify "Effect": "Deny" in the second SCP to override the "Effect": "Allow"
// in the FullAWSAccess policy (or any other attached SCP), you're using the
// authorization strategy of a "[deny list] ".
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [deny list]: https://docs.aws.amazon.com/organizations/latest/userguide/SCP_strategies.html#orgs_policies_denylist
// [allow list]: https://docs.aws.amazon.com/organizations/latest/userguide/SCP_strategies.html#orgs_policies_allowlist
