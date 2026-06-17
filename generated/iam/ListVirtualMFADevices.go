package iam

// ListVirtualMFADevices is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists the virtual MFA devices defined in the Amazon Web Services account by
// assignment status. If you do not specify an assignment status, the operation
// returns a list of all virtual MFA devices. Assignment status can be Assigned ,
// Unassigned , or Any .
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view tag information for a virtual
// MFA device, see [ListMFADeviceTags].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [ListMFADeviceTags]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListMFADeviceTags.html
