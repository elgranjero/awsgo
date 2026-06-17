package storagegateway

// JoinDomain is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Adds a file gateway to an Active Directory domain. This operation is only
// supported for file gateways that support the SMB file protocol.
//
// Joining a domain creates an Active Directory computer account in the default
// organizational unit, using the gateway's Gateway ID as the account name (for
// example, SGW-1234ADE). If your Active Directory environment requires that you
// pre-stage accounts to facilitate the join domain process, you will need to
// create this account ahead of time.
//
// To create the gateway's computer account in an organizational unit other than
// the default, you must specify the organizational unit when joining the domain.
