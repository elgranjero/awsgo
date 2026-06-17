package directconnect

// AllocatePublicVirtualInterface is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Provisions a public virtual interface to be owned by the specified Amazon Web
// Services account.
//
// The owner of a connection calls this function to provision a public virtual
// interface to be owned by the specified Amazon Web Services account.
//
// Virtual interfaces created using this function must be confirmed by the owner
// using ConfirmPublicVirtualInterface. Until this step has been completed, the virtual interface is in the
// confirming state and is not available to handle traffic.
//
// When creating an IPv6 public virtual interface, omit the Amazon address and
// customer address. IPv6 addresses are automatically assigned from the Amazon pool
// of IPv6 addresses; you cannot specify custom IPv6 addresses.
