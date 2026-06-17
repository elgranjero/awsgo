package globalaccelerator

// DeleteAccelerator is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Delete an accelerator. Before you can delete an accelerator, you must disable
// it and remove all dependent resources (listeners and endpoint groups). To
// disable the accelerator, update the accelerator to set Enabled to false.
//
// When you create an accelerator, by default, Global Accelerator provides you
// with a set of two static IP addresses. Alternatively, you can bring your own IP
// address ranges to Global Accelerator and assign IP addresses from those ranges.
//
// The IP addresses are assigned to your accelerator for as long as it exists,
// even if you disable the accelerator and it no longer accepts or routes traffic.
// However, when you delete an accelerator, you lose the static IP addresses that
// are assigned to the accelerator, so you can no longer route traffic by using
// them. As a best practice, ensure that you have permissions in place to avoid
// inadvertently deleting accelerators. You can use IAM policies with Global
// Accelerator to limit the users who have permissions to delete an accelerator.
// For more information, see [Identity and access management]in the Global Accelerator Developer Guide.
//
// [Identity and access management]: https://docs.aws.amazon.com/global-accelerator/latest/dg/auth-and-access-control.html
