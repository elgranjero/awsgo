package globalaccelerator

// CreateCrossAccountAttachment is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Create a cross-account attachment in Global Accelerator. You create a
// cross-account attachment to specify the principals who have permission to work
// with resources in accelerators in their own account. You specify, in the same
// attachment, the resources that are shared.
//
// A principal can be an Amazon Web Services account number or the Amazon Resource
// Name (ARN) for an accelerator. For account numbers that are listed as
// principals, to work with a resource listed in the attachment, you must sign in
// to an account specified as a principal. Then, you can work with resources that
// are listed, with any of your accelerators. If an accelerator ARN is listed in
// the cross-account attachment as a principal, anyone with permission to make
// updates to the accelerator can work with resources that are listed in the
// attachment.
//
// Specify each principal and resource separately. To specify two CIDR address
// pools, list them individually under Resources , and so on. For a command line
// operation, for example, you might use a statement like the following:
//
// "Resources": [{"Cidr": "169.254.60.0/24"},{"Cidr": "169.254.59.0/24"}]
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
