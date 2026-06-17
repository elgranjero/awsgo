package shield

// AssociateDRTRole is generated as a reference stub.
// Executable command wiring lives under cmd/shield.go.
//
// Authorizes the Shield Response Team (SRT) using the specified role, to access
// your Amazon Web Services account to assist with DDoS attack mitigation during
// potential attacks. This enables the SRT to inspect your WAF configuration and
// create or update WAF rules and web ACLs.
//
// You can associate only one RoleArn with your subscription. If you submit an
// AssociateDRTRole request for an account that already has an associated role, the
// new RoleArn will replace the existing RoleArn .
//
// Prior to making the AssociateDRTRole request, you must attach the
// AWSShieldDRTAccessPolicy managed policy to the role that you'll specify in the
// request. You can access this policy in the IAM console at [AWSShieldDRTAccessPolicy]. For more
// information see [Adding and removing IAM identity permissions]. The role must also trust the service principal
// drt.shield.amazonaws.com . For more information, see [IAM JSON policy elements: Principal].
//
// The SRT will have access only to your WAF and Shield resources. By submitting
// this request, you authorize the SRT to inspect your WAF and Shield configuration
// and create and update WAF rules and web ACLs on your behalf. The SRT takes these
// actions only if explicitly authorized by you.
//
// You must have the iam:PassRole permission to make an AssociateDRTRole request.
// For more information, see [Granting a user permissions to pass a role to an Amazon Web Services service].
//
// To use the services of the SRT and make an AssociateDRTRole request, you must
// be subscribed to the [Business Support plan]or the [Enterprise Support plan].
//
// [Adding and removing IAM identity permissions]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html
// [Enterprise Support plan]: http://aws.amazon.com/premiumsupport/enterprise-support/
// [AWSShieldDRTAccessPolicy]: https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy
// [IAM JSON policy elements: Principal]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html
// [Granting a user permissions to pass a role to an Amazon Web Services service]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [Business Support plan]: http://aws.amazon.com/premiumsupport/business-support/
