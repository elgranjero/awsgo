package globalaccelerator

// DeprovisionByoipCidr is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Releases the specified address range that you provisioned to use with your
// Amazon Web Services resources through bring your own IP addresses (BYOIP) and
// deletes the corresponding address pool.
//
// Before you can release an address range, you must stop advertising it by using [WithdrawByoipCidr]
// and you must not have any accelerators that are using static IP addresses
// allocated from its address range.
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
// [WithdrawByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/WithdrawByoipCidr.html
