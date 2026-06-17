package workspaces

// ListAvailableManagementCidrRanges is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Retrieves a list of IP address ranges, specified as IPv4 CIDR blocks, that you
// can use for the network management interface when you enable Bring Your Own
// License (BYOL).
//
// This operation can be run only by Amazon Web Services accounts that are enabled
// for BYOL. If your account isn't enabled for BYOL, you'll receive an
// AccessDeniedException error.
//
// The management network interface is connected to a secure Amazon WorkSpaces
// management network. It is used for interactive streaming of the WorkSpace
// desktop to Amazon WorkSpaces clients, and to allow Amazon WorkSpaces to manage
// the WorkSpace.
