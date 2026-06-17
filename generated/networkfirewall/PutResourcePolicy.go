package networkfirewall

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
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
