package storagegateway

// DescribeWorkingStorage is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Returns information about the working storage of a gateway. This operation is
// only supported in the stored volumes gateway type. This operation is deprecated
// in cached volumes API version (20120630). Use DescribeUploadBuffer instead.
//
// Working storage is also referred to as upload buffer. You can also use the
// DescribeUploadBuffer operation to add upload buffer to a stored volume gateway.
//
// The response includes disk IDs that are configured as working storage, and it
// includes the amount of working storage allocated and used.
