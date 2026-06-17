package iam

// CreateVirtualMFADevice is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Creates a new virtual MFA device for the Amazon Web Services account. After
// creating the virtual MFA, use [EnableMFADevice]to attach the MFA device to an IAM user. For more
// information about creating and working with virtual MFA devices, see [Using a virtual MFA device]in the IAM
// User Guide.
//
// For information about the maximum number of MFA devices you can create, see [IAM and STS quotas] in
// the IAM User Guide.
//
// The seed information contained in the QR code and the Base32 string should be
// treated like any other secret access information. In other words, protect the
// seed information as you would your Amazon Web Services access keys or your
// passwords. After you provision your virtual device, you should ensure that the
// information is destroyed following secure procedures.
//
// [Using a virtual MFA device]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html
// [EnableMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_EnableMFADevice.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
