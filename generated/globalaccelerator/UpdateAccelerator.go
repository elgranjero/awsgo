package globalaccelerator

// UpdateAccelerator is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Update an accelerator to make changes, such as the following:
//
// - Change the name of the accelerator.
//
// - Disable the accelerator so that it no longer accepts or routes traffic, or
// so that you can delete it.
//
// - Enable the accelerator, if it is disabled.
//
// - Change the IP address type to dual-stack if it is IPv4, or change the IP
// address type to IPv4 if it's dual-stack.
//
// Be aware that static IP addresses remain assigned to your accelerator for as
// long as it exists, even if you disable the accelerator and it no longer accepts
// or routes traffic. However, when you delete the accelerator, you lose the static
// IP addresses that are assigned to it, so you can no longer route traffic by
// using them.
//
// Global Accelerator is a global service that supports endpoints in multiple
// Amazon Web Services Regions but you must specify the US West (Oregon) Region to
// create, update, or otherwise work with accelerators. That is, for example,
// specify --region us-west-2 on Amazon Web Services CLI commands.
